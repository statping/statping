/// <reference types="cypress" />

import "../support/commands"

context('Groups Tests', () => {


  beforeEach(() => {
    cy.restoreLocalStorageCache();
  });

  afterEach(() => {
    cy.saveLocalStorageCache();
  });

  it('should Login', () => {
    cy.visit('/login')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('.navbar-brand').should('contain', 'Statping')
    cy.getCookies()

    cy.getCookies().should('have.length', 1)
  })

  it('should goto groups', () => {
    cy.visit('/dashboard/services')
    cy.get('.sortable_groups > tr').should('have.length', 3)

    cy.get('.sortable_groups > tr').eq(0).contains('PUBLIC')
    cy.get('.sortable_groups > tr').eq(1).contains('PUBLIC')
    cy.get('.sortable_groups > tr').eq(2).contains('PRIVATE')
  })

  it('should create new Group', () => {
    cy.visit('/dashboard/services')
    cy.get('#title').clear().type('Test Group')

    cy.get('button[type="submit"]').click()
  })

  it('should create new Private Group', () => {
    cy.visit('/dashboard/services')
    cy.get('#title').clear().type('Test Private Group')
    cy.get('.float-left > label').click()
    cy.get('button[type="submit"]').click()
  })

  it('should edit Group', () => {
    cy.visit('/dashboard/services')
    cy.get('.sortable_groups > tr').eq(0).find('.btn-outline-secondary').click()
    cy.get('#title').should('have.value', 'Test Group')
    cy.get('#title').clear().type('Updated Group')
    cy.get('button[type="submit"]').click()

    cy.get('.sortable_groups > tr').eq(0).contains('Updated Group')
  })

  it('should confirm new groups', () => {
    cy.visit('/dashboard/services')
    cy.get('.sortable_groups > tr').should('have.length', 5)

    cy.get('.sortable_groups > tr').eq(0).contains('PUBLIC')
    cy.get('.sortable_groups > tr').eq(0).contains('Updated Group')
    cy.get('.sortable_groups > tr').eq(1).contains('PRIVATE')
    cy.get('.sortable_groups > tr').eq(1).contains('Test Private Group')
  })

    it('should delete new groups', () => {
      cy.visit('/dashboard/services')
        cy.get('.sortable_groups > tr').eq(0).find('.btn-danger').click()
        cy.get('.sortable_groups > tr').eq(1).find('.btn-danger').click()
        cy.get('.sortable_groups > tr').should('have.length', 3)
    })

})
